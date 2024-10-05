import 'package:dio/dio.dart';

import 'tokens.dart';

// Set default configs
Dio configureClient(TokenHandler tokenHandler) {
  Dio dio = Dio(BaseOptions(
    baseUrl: 'https://localhost:8000',
    headers: {
      'Accept': 'application/json',
      'Content-Type': 'application/json',
    },
  ));

  dio.interceptors.add(InterceptorsWrapper(
    onRequest: (options, handler) {
      final token = tokenHandler.getAccessToken(); // get token from local storage
      options.headers['Authorization'] = 'Bearer $token';
      return handler.next(options); // continue
    },
    onError: (DioException e, handler) async {
      // Check if the error is due to an expired token
      if (e.response?.statusCode == 401) {
      // Attempt to refresh the token
        try {
          final newToken = await tokenHandler.refreshToken();
          // Update the request with the new token
          final options = e.requestOptions;
          options.headers['Authorization'] = 'Bearer $newToken';
          // Retry the request with the new token
          final cloneReq = await dio.request(
            options.path,
            options: Options(
              method: options.method,
              headers: options.headers,
            ),
            data: options.data,
            queryParameters: options.queryParameters,
          );
          return handler.resolve(cloneReq);
        } catch (refreshError) {
          // If refresh fails, forward the original error
          return handler.next(e);
        }
      }
      // If the error is not due to an expired token, forward the error
      return handler.next(e); // continue
    },
  ));

  return dio;
}

import 'package:dio/dio.dart';
import 'package:shared_preferences/shared_preferences.dart';

import 'errors/exceptions.dart';

abstract class TokenHandler {
  Future<void> setAccessToken(String token);
  Future<void> setRefreshToken(String token);
  Future<String> getAccessToken();
  Future<String> getRefreshToken();
  Future<String> refreshToken();
  Future<void> removeToken();
}

class TokenHandlerImpl implements TokenHandler {
  final SharedPreferences sharedPreferences;

  TokenHandlerImpl({required this.sharedPreferences});

  @override
  Future<void> removeToken() async {
    if (sharedPreferences.containsKey('token')) {
      await sharedPreferences.remove('token');
    } else {
      throw CacheException();
    }
  }
  
  @override
  Future<String> getAccessToken() async {
    return sharedPreferences.getString('token') ?? '';
  }
  
  @override
  Future<String> getRefreshToken() async {
    return sharedPreferences.getString('refresh_token') ?? '';
  }
  
  @override
  Future<void> setAccessToken(String token) async {
    await sharedPreferences.setString('token', token);
  }
  
  @override
  Future<void> setRefreshToken(String token) async {
    await sharedPreferences.setString('refresh_token', token);
  }
  
  @override
  Future<String> refreshToken() async {
    final refreshToken = await getRefreshToken();
    if (refreshToken.isEmpty) {
      throw CacheException();
    }
    final response = await Dio().post('http://localhost:8000/refresh', data: {'refresher': refreshToken});
    if (response.statusCode != 200) {
      throw ServerException();
    }
    final String newAccessToken = response.data['access_token'];
    await setAccessToken(newAccessToken);
    return newAccessToken;
    
  }
}
import 'package:dio/dio.dart';

import '../../../../core/errors/exceptions.dart';
import '../models/tokens_model.dart';
import '../models/user_model.dart';

abstract class RemoteDatasource {
  Future<Token> login(String email, String password);
  Future<void> register(UserModel user);
}

class RemoteDatasourceImpl implements RemoteDatasource {
  final Dio client;
  RemoteDatasourceImpl(this.client);
  
  @override
  Future<Token> login(String email, String password) async {
    final response = await client.post('/login', data: {
      'email': email,
      'password': password,
    },);
    if (response.statusCode != 200) {
      throw ServerException();
    }
    return Token.fromJson(response.data);
  }

  @override
  Future<void> register(UserModel user) async {
    final response = await client.post('/register', data: user.toJson());
    if (response.statusCode != 200) {
      throw ServerException();
    }
  }
}
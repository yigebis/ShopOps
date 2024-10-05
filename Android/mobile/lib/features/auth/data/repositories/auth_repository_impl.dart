import '../../../../core/errors/exceptions.dart';
import '../../../../core/errors/failure.dart';
import '../../../../core/tokens.dart';
import '../../domain/entities/user.dart';
import 'package:dartz/dartz.dart';

import '../../domain/repositories/repository.dart';
import '../datasources/remote_datasource.dart';
import '../models/tokens_model.dart';
import '../models/user_model.dart';

class AuthRepositoryImpl extends AuthRepository {
  final RemoteDatasource remoteDatasource;
  final TokenHandler  tokenHandler;

  AuthRepositoryImpl({required this.remoteDatasource, required this.tokenHandler});

  @override
  Future<Either<Failure, void>> logIn({required String email, required String password}) async {
    try {
      final Token token = await remoteDatasource.login(email, password);
      await tokenHandler.setAccessToken(token.accessToken);
      await tokenHandler.setRefreshToken(token.refreshToken);
      // store the tokens in the local storage
      return const Right(null);
    } on ServerException {
      return const Left(ServerFailure('Server Failure'));
    } catch (e) {
      return const Left(UnexpectedFailure('Unexpected Failure'));
    }
  }

  @override
  Future<Either<Failure, void>> signUp({required User user}) async {
    try {
      await remoteDatasource.register(UserModel.fromEntity(user));
      return const Right(null);
    } on ServerException {
      return const Left(ServerFailure('Server Failure'));
    } catch (e) {
      return const Left(UnexpectedFailure('Unexpected Failure'));
    }
  }
}
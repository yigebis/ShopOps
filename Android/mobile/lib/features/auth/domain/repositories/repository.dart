import 'package:dartz/dartz.dart';

import '../../../../core/errors/failure.dart';
import '../entities/user.dart';

abstract class AuthRepository {
  Future<Either<Failure, void>> logIn({required String email, required String password});
  Future<Either<Failure, void>> signUp({required User user});
}
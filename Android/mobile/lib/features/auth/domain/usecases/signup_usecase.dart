import 'package:dartz/dartz.dart';
import 'package:equatable/equatable.dart';

import '../../../../core/errors/failure.dart';
import '../../../../core/usecase.dart';
import '../entities/user.dart';
import '../repositories/repository.dart';

class SignupUsecase extends UseCase<void, SignUpParams> {
  final AuthRepository _repository;

  SignupUsecase(this._repository);

  @override
  Future<Either<Failure, void>> call(SignUpParams params) async {
    return await _repository.signUp(user: params.user);
  }
}

class SignUpParams extends Equatable {
  final User user;

  const SignUpParams({required this.user});

  @override
  List<Object?> get props => [user];
}

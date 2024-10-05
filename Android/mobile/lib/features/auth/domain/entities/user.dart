import 'package:equatable/equatable.dart';

class User extends Equatable {
  final String? id;
  final String email;
  final String firstname;
  final String lastname;
  final String phonenumber;
  final String sex;
  final String? password;

  const User({
    this.id,
    required this.email,
    required this.firstname,
    required this.lastname,
    required this.phonenumber,
    required this.sex,
    this.password,
  });

  @override
  List<Object?> get props => [id, email, firstname, lastname, phonenumber,  sex, password];
}
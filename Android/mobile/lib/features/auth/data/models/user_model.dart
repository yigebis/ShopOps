import '../../domain/entities/user.dart';

class UserModel extends User {
  const UserModel({
    super.id,
    required super.email,
    required super.firstname,
    required super.lastname,
    super.password,
    required super.phonenumber, 
    required super.sex
  });

  factory UserModel.fromJson(Map<String, dynamic> json) {
    return UserModel(
      id: json['id'],
      email: json['email'],
      firstname: json['first_name'],
      lastname: json['last_name'],
      sex: json['sex'],
      phonenumber: json['phone_number'],
    );
  }

  factory UserModel.fromEntity(User user) {
    return UserModel(
      email: user.email,
      firstname: user.firstname,
      lastname: user.lastname,
      sex: user.sex,
      phonenumber: user.phonenumber,
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'email': email,
      'firs_nme': firstname,
      'last_name': lastname,
      'password': password,
      'phone_number': phonenumber,
      'sex': sex,
    };
  }
}
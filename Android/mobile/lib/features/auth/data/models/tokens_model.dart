import 'package:equatable/equatable.dart';

class Token extends Equatable {
  final String accessToken;
  final String refreshToken;

  const Token({
    required this.accessToken,
    required this.refreshToken,
  });

  factory Token.fromJson(Map<String, dynamic> json) {
    return Token(
      accessToken: json['token'],
      refreshToken: json['refresher'],
    );
  }

  @override
  List<Object?> get props => [accessToken, refreshToken];
}
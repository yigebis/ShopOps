import 'package:flutter/material.dart';

import 'colors.dart';

final ThemeData appTheme = _buildAppTheme();

ThemeData _buildAppTheme() {
  final ThemeData base = ThemeData.light(useMaterial3: true);
  return base.copyWith(
    colorScheme: base.colorScheme.copyWith(
      primary: primaryGreen,
      onPrimary: Colors.white,
      secondary: secondaryGreen,
      error: errorRed,
    ),
    textTheme: _buildShrineTextTheme(base.textTheme),
    textSelectionTheme: const TextSelectionThemeData(
      selectionColor: surfaceGreen,
    ),
    // TODO: Decorate the inputs (103)
  );
}

TextTheme _buildShrineTextTheme(TextTheme base) {
  return base
    .copyWith(
      headlineSmall: base.headlineSmall!.copyWith(
        fontWeight: FontWeight.w500,
      ),
      titleLarge: base.titleLarge!.copyWith(
        fontSize: 18.0,
      ),
      bodySmall: base.bodySmall!.copyWith(
        fontWeight: FontWeight.w400,
        fontSize: 14.0,
      ),
      bodyLarge: base.bodyLarge!.copyWith(
        fontWeight: FontWeight.w500,
        fontSize: 16.0,
      ),
    );
}
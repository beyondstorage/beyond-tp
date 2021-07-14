import 'package:flutter/material.dart';

import '../common/colors.dart';

TextStyle getTextStyle({
  Color? color,
  double? fontSize,
  String? fontFamily,
  FontWeight? fontWeight,
  TextDecoration? decoration,
  double? height,
}) =>
    TextStyle(
      height: height ?? 1.67,
      fontFamily: "Roboto",
      fontSize: fontSize ?? 12.0,
      fontWeight: fontWeight ?? FontWeight.normal,
      color: color ?? regularFontColor,
      decoration: decoration ?? TextDecoration.none,
    );

TextTheme getTextTheme({Color? color}) => TextTheme(
      headline1: getTextStyle(
        color: color ?? headlineFontColor,
        height: 1.0,
        fontSize: 48,
        fontWeight: FontWeight.w600,
      ),
      headline2: getTextStyle(
        color: color ?? headlineFontColor,
        height: 1.0,
        fontSize: 24,
        fontWeight: FontWeight.w600,
      ),
      headline3: getTextStyle(
        color: color ?? headlineFontColor,
        height: 1.0,
        fontSize: 20,
        fontWeight: FontWeight.w600,
      ),
      headline4: getTextStyle(
        color: color ?? headlineFontColor,
        height: 1.0,
        fontSize: 18,
        fontWeight: FontWeight.w600,
      ),
      headline5: getTextStyle(
        color: color ?? headlineFontColor,
        height: 1.50,
        fontSize: 16,
        fontWeight: FontWeight.w400,
      ),
      headline6: getTextStyle(
        color: color ?? headlineFontColor,
        fontSize: 14,
        fontWeight: FontWeight.w600,
      ),
      bodyText1: getTextStyle(
        color: color ?? regularFontColor,
      ),
      bodyText2: getTextStyle(
        color: color ?? secondaryFontColor,
      ),
    );

final themeData = ThemeData(
  brightness: Brightness.light,
  fontFamily: "Roboto",
  primaryColor: primaryColor,
  primaryColorLight: rgba(50, 69, 88, 1),
  backgroundColor: rgba(255, 255, 255, 1),
  scaffoldBackgroundColor: rgba(231, 238, 242, 1),
  errorColor: rgba(202, 38, 33, 1),
  appBarTheme: AppBarTheme(
    backgroundColor: rgba(2, 5, 8, 1),
    textTheme: getTextTheme(color: rgba(255, 255, 255, 1)),
  ),
  primaryTextTheme: TextTheme(
    headline6: getTextStyle(
        color: rgba(0, 170, 114, 1), fontSize: 14, fontWeight: FontWeight.w600),
  ),
  textTheme: getTextTheme(),
  dataTableTheme: DataTableThemeData(
      dataRowHeight: 44.0,
      dataTextStyle: getTextStyle(),
      headingRowHeight: 44.0,
      headingRowColor: MaterialStateProperty.resolveWith<Color>(
          (Set<MaterialState> states) => rgba(245, 247, 250, 1)),
      headingTextStyle: getTextStyle(
        fontWeight: FontWeight.w600,
      ),
      decoration: BoxDecoration()),
);

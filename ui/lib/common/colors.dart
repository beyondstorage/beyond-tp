import 'package:flutter/material.dart';

double opacity = 0.5;

Color rgba(int r, int g, int b, double a) => Color.fromRGBO(r, g, b, a);

Color headlineFontColor = rgba(100, 116, 139, 1);
Color regularFontColor = rgba(71, 85, 105, 1);
Color secondaryFontColor = rgba(100, 116, 139, 1);
Color disableFontColor = rgba(148, 163, 184, 1);

Color defaultColor = rgba(203, 213, 225, 1);
Color defaultDisabledColor = rgba(226, 232, 240, 1);

Color primaryColor = rgba(59, 130, 246, 1);
Color primaryHoveredColor = rgba(37, 99, 235, 1);
Color primaryPressedColor = rgba(29, 78, 216, 1);
Color primaryDisabledColor = rgba(59, 130, 246, opacity);
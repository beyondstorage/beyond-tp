import 'package:flutter/material.dart';

import './constants.dart';
import '../../common/colors.dart';

Color defaultColor = rgba(203, 213, 225, 1);
Color defaultDisabledColor = rgba(226, 232, 240, 1);

Color primaryColor = rgba(59, 130, 246, 1);
Color primaryHoveredColor = rgba(37, 99, 235, 1);
Color primaryPressedColor = rgba(29, 78, 216, 1);
Color primaryDisabledColor = rgba(59, 130, 246, opacity);

Color getPrimaryBackgroundColor(ButtonType type) {
  switch (type) {
    case ButtonType.error:
      return rgba(202, 38, 33, 1);
    case ButtonType.primary:
      return primaryColor;
    default:
      return Colors.white;
  }
}

Color getDisabledBackgroundColor(ButtonType type) {
  switch (type) {
    case ButtonType.error:
      return rgba(202, 38, 33, opacity);
    case ButtonType.primary:
      return primaryDisabledColor;
    default:
      return Colors.white;
  }
}

Color getHoveredBackgroundColor(ButtonType type) {
  switch (type) {
    case ButtonType.error:
      return rgba(202, 38, 33, 1);
    case ButtonType.primary:
      return primaryHoveredColor;
    default:
      return Colors.white;
  }
}

Color getPressedBackgroundColor(ButtonType type) {
  switch (type) {
    case ButtonType.error:
      return rgba(202, 38, 33, 1);
    case ButtonType.primary:
      return primaryPressedColor;
    default:
      return Colors.white;
  }
}

Color getDefaultFontColor(Set<MaterialState> states, bool disabled) {
  if (disabled) {
    return disableFontColor;
  }

  if (states.contains(MaterialState.pressed)) {
    return primaryPressedColor;
  }

  if (states.contains(MaterialState.hovered)) {
    return primaryHoveredColor;
  }

  return regularFontColor;
}

BorderSide getDefaultOutLineBorderSide(Set<MaterialState> states, bool disabled) {
  if (disabled) {
    return BorderSide(color: defaultDisabledColor);
  }

  Color color = defaultColor;

  if (states.contains(MaterialState.pressed)) {
    color = primaryPressedColor;
  } else if (states.contains(MaterialState.hovered)) {
    color = primaryHoveredColor;
  }

  return BorderSide(color: color);
}


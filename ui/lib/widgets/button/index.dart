import 'dart:ui';
import 'package:flutter/material.dart';
import 'package:flutter/rendering.dart';

import './colors.dart';
import './constants.dart';

class Button extends StatelessWidget {
  final IconData? icon;
  final Widget child;
  final VoidCallback onPressed;
  final ButtonType type;
  final bool disabled;

  Button({
    this.icon,
    required this.child,
    required this.onPressed,
    this.disabled = false,
    this.type = ButtonType.defaults,
  });

  Color getFontColor(Set<MaterialState> states) {
    if (type == ButtonType.defaults) {
      return getDefaultFontColor(states, disabled);
    }

    return Colors.white;
  }

  Color getBackgroundColor(Set<MaterialState> states) {
    if (disabled) {
      return getDisabledBackgroundColor(type);
    }

    if (states.contains(MaterialState.pressed)) {
      return getPressedBackgroundColor(type);
    }

    if (states.contains(MaterialState.hovered)) {
      return getHoveredBackgroundColor(type);
    }

    return getPrimaryBackgroundColor(type);
  }

  BorderSide getSide(Set<MaterialState> states) {
    if (type == ButtonType.defaults) {
      return getDefaultOutLineBorderSide(states, disabled);
    }

    return BorderSide(color: getBackgroundColor(states));
  }

  OutlinedBorder getShape(Set<MaterialState> states) {
    return RoundedRectangleBorder(borderRadius: BorderRadius.circular(30.0));
  }

  MouseCursor getMouseCursor(Set<MaterialState> states) {
    return disabled ? SystemMouseCursors.forbidden : SystemMouseCursors.click;
  }

  Size getSize(Set<MaterialState> states) => Size(0, 0);

  TextStyle getTextStyle(Set<MaterialState> states) {
    return TextStyle(fontSize: 12);
  }

  void onClick() {
    if (!disabled) onPressed();
  }

  @override
  Widget build(BuildContext context) {
    if (icon == null) {
      return OutlinedButton(
        onPressed: onClick,
        style: ButtonStyle(
          padding: MaterialStateProperty.all(EdgeInsets.zero),
          foregroundColor: MaterialStateProperty.resolveWith(getFontColor),
          backgroundColor:
              MaterialStateProperty.resolveWith(getBackgroundColor),
          side: MaterialStateProperty.resolveWith(getSide),
          shape: MaterialStateProperty.resolveWith(getShape),
          textStyle: MaterialStateProperty.resolveWith(getTextStyle),
          minimumSize: MaterialStateProperty.resolveWith(getSize),
          mouseCursor: MaterialStateProperty.resolveWith(getMouseCursor),
        ),
        child: Container(
          height: 32,
          alignment: Alignment.center,
          padding: EdgeInsets.symmetric(horizontal: 24.0),
          child: child,
        ),
      );
    }

    return OutlinedButton.icon(
      onPressed: onClick,
      style: ButtonStyle(
        alignment: Alignment.center,
        padding:
            MaterialStateProperty.all(EdgeInsets.symmetric(horizontal: 24)),
        foregroundColor: MaterialStateProperty.resolveWith(getFontColor),
        backgroundColor: MaterialStateProperty.resolveWith(getBackgroundColor),
        side: MaterialStateProperty.resolveWith(getSide),
        shape: MaterialStateProperty.resolveWith(getShape),
        textStyle: MaterialStateProperty.resolveWith(getTextStyle),
        minimumSize: MaterialStateProperty.resolveWith(getSize),
        mouseCursor: MaterialStateProperty.resolveWith(getMouseCursor),
      ),
      icon: Icon(icon, size: 14),
      label: Container(
        height: 32,
        alignment: Alignment.center,
        child: child,
      ),
    );
  }
}

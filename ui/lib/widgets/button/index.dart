import 'package:flutter/material.dart';

import '../../common/colors.dart';

enum ButtonType {
  error,
  primary,
  defaults,
}

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
    double opacity = disabled ? 0.5 : 1.00;
    Color color;

    switch (type) {
      case ButtonType.error:
        color = rgba(202, 38, 33, 1);
        break;
      case ButtonType.primary:
        color = Colors.white;
        break;
      default:
        color = rgba(50, 69, 88, opacity);
    }

    return color;
  }

  Color getBackgroundColor(Set<MaterialState> states) {
    double opacity = disabled ? 0.5 : 1.00;

    switch (type) {
      case ButtonType.error:
        return rgba(202, 38, 33, 1);
      case ButtonType.primary:
        return rgba(0, 170, 114, opacity);
      default:
        return Colors.white;
    }
  }

  BorderSide getSide(Set<MaterialState> states) {
    double opacity = disabled ? 0.5 : 1.00;
    Color color;

    switch (type) {
      case ButtonType.error:
        color = rgba(202, 38, 33, 1);
        break;
      case ButtonType.primary:
        color = rgba(0, 170, 114, opacity);
        break;
      default:
        color = rgba(140, 140, 140, opacity);
    }

    return BorderSide(
      style: BorderStyle.solid,
      color: color,
    );
  }

  OutlinedBorder getShape(Set<MaterialState> states) {
    return RoundedRectangleBorder(borderRadius: BorderRadius.circular(30.0));
  }

  Size getSize(Set<MaterialState> states) => Size(0, 0);

  TextStyle getTextStyle(Set<MaterialState> states) {
    return TextStyle(fontSize: 12);
  }

  @override
  Widget build(BuildContext context) {
    if (icon == null) {
      return OutlinedButton(
        onPressed: onPressed,
        style: ButtonStyle(
          padding: MaterialStateProperty.all(EdgeInsets.zero),
          foregroundColor: MaterialStateProperty.resolveWith(getFontColor),
          backgroundColor:
              MaterialStateProperty.resolveWith(getBackgroundColor),
          side: MaterialStateProperty.resolveWith(getSide),
          shape: MaterialStateProperty.resolveWith(getShape),
          textStyle: MaterialStateProperty.resolveWith(getTextStyle),
          minimumSize: MaterialStateProperty.resolveWith(getSize),
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
      onPressed: onPressed,
      style: ButtonStyle(
        padding:
            MaterialStateProperty.all(EdgeInsets.symmetric(horizontal: 24)),
        alignment: Alignment.center,
        foregroundColor: MaterialStateProperty.resolveWith(getFontColor),
        backgroundColor: MaterialStateProperty.resolveWith(getBackgroundColor),
        side: MaterialStateProperty.resolveWith(getSide),
        shape: MaterialStateProperty.resolveWith(getShape),
        textStyle: MaterialStateProperty.resolveWith(getTextStyle),
        minimumSize: MaterialStateProperty.resolveWith(getSize),
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

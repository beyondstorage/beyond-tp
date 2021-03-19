import 'package:flutter/material.dart';

import '../../common/colors.dart';

enum ButtonType {
  error,
  primary,
  defaults,
}

class Button extends StatelessWidget{
  final IconData icon;
  final Widget label;
  final Function onPressed;
  final ButtonType type;
  final bool disabled;

  Button({
    this.icon,
    this.label,
    this.onPressed,
    this.disabled = false,
    this.type = ButtonType.defaults
  });

  Color getColor(Set<MaterialState> states) {
    // const Set<MaterialState> interactiveStates = <MaterialState>{
    //   MaterialState.pressed,
    //   MaterialState.hovered,
    //   MaterialState.focused,
    // };

    double opacity = disabled ? 0.5 : 1.00;

    switch(type) {
      case ButtonType.error:
        return Colors.red;
      case ButtonType.primary:
        return rgba(0, 170, 114, opacity);
      default:
        return Colors.grey;
    }
  }

  EdgeInsetsGeometry getPadding(Set<MaterialState> states) {
    return EdgeInsets.symmetric(horizontal: 24.0, vertical: 15.0);
  }

  TextStyle getTextStyle(Set<MaterialState> states) {
    return TextStyle(fontSize: 12);
  }

  Size getSize(Set<MaterialState> states) => Size(90.0, 32.0);

  @override
  Widget build(BuildContext context) {
    return ElevatedButton.icon(
      label: label,
      onPressed: onPressed,
      icon: Icon(icon, size: 14),
      style: ButtonStyle(
        padding: MaterialStateProperty.resolveWith(getPadding),
        minimumSize: MaterialStateProperty.resolveWith(getSize),
        textStyle: MaterialStateProperty.resolveWith(getTextStyle),
        backgroundColor: MaterialStateProperty.resolveWith(getColor),
      ),
    );
  }
}
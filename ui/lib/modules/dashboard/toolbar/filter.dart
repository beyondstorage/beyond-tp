import 'package:flutter/material.dart';

import '../../../common/colors.dart';


class Filter extends StatelessWidget {
  final int? counts;
  final String title;
  final VoidCallback onPressed;
  final bool selected;

  Filter({
    this.counts = 0,
    required this.selected,
    required this.title,
    required this.onPressed,
  });

  Color getForeGroundColor(Set<MaterialState> states) {
    if (states.contains(MaterialState.hovered) || selected) {
      return primaryColor;
    }

    return regularFontColor;
  }

  @override
  Widget build(BuildContext context) {
    return Container(
      margin: EdgeInsets.only(right: 20.0),
      padding: EdgeInsets.only(bottom: 8.0),
      decoration: BoxDecoration(
        color: Colors.white,
        boxShadow: selected ? [
          BoxShadow(offset: Offset(0, 1), color: primaryColor),
          BoxShadow(offset: Offset(-1, 0), color: Colors.white),
          BoxShadow(offset: Offset(1, 0), color: Colors.white),
        ] : [],
      ),
      child: TextButton(
        style: ButtonStyle(
          foregroundColor: MaterialStateProperty.resolveWith(getForeGroundColor),
          overlayColor: MaterialStateProperty.resolveWith((states) => Colors.white),
        ),
        onPressed: onPressed,
        child: Text("$title($counts)"),
      ),
    );
  }
}
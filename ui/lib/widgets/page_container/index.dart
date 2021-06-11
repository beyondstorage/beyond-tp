import 'package:flutter/material.dart';

import '../../common/colors.dart';

class WidgetContainer extends StatelessWidget {
  final Widget child;
  final EdgeInsetsGeometry? margin;

  WidgetContainer({
    this.margin,
    required this.child,
  });

  @override
  Widget build(BuildContext context) {
    return Container(
      margin: margin,
      decoration: new BoxDecoration(
        color: Colors.white,
        borderRadius: BorderRadius.all(Radius.circular(6.0)),
        boxShadow: [
          BoxShadow(
            offset: Offset(0, 1),
            color: rgba(52, 61, 190, 0.1),
            blurRadius: 3.0,
          )
        ],
      ),
      child: child,
    );
  }
}


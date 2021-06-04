import 'package:flutter/material.dart';

import '../../common/colors.dart';

enum ButtonType {
  error,
  primary,
  defaults,
}

class Toggle extends StatelessWidget {
  final bool value;
  final Function onChange;

  Toggle({
    required this.value,
    required this.onChange,
  });

  Color getColor(BuildContext context) {
    if (value == true) {
      return Theme.of(context).primaryColor;
    }

    return rgba(213, 222, 231, 1);
  }

  @override
  Widget build(BuildContext context) {
    return Listener(
        onPointerUp: (PointerUpEvent event) => onChange!(value != true),
        child: Container(
          width: 28,
          height: 16,
          padding: EdgeInsets.symmetric(horizontal: 2),
          alignment:
              value == true ? Alignment.centerRight : Alignment.centerLeft,
          decoration: BoxDecoration(
            color: getColor(context),
            borderRadius: BorderRadius.all(Radius.circular(16)),
          ),
          child: Container(
            width: 12,
            height: 12,
            decoration: BoxDecoration(
              color: Colors.white,
              borderRadius: BorderRadius.all(Radius.circular(12)),
            ),
          ),
        ));
  }
}

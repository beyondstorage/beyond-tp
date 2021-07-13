import 'package:flutter/material.dart';
import 'package:flutter/rendering.dart';

import '../../common/colors.dart';
import 'model.dart';

class RadioGroup extends StatelessWidget {
  final String? value;
  final List<RadioOption> options;
  final void Function(String?)? onChange;

  RadioGroup({
    this.value,
    required this.options,
    this.onChange,
  });

  List<Widget> getChildren() {
    List<Widget> children = [];

    options.forEach((option) {
      children.addAll([
        GestureDetector(
          onTap: () => print(option.value),
          child: MouseRegion(
            cursor: SystemMouseCursors.click,
            child: Row(
              children: [
                Radio(
                  value: option.value,
                  groupValue: value,
                  onChanged: (v) {},
                  splashRadius: 0,
                  focusColor: primaryHoveredColor,
                  hoverColor: primaryHoveredColor,
                  activeColor: primaryColor,
                ),
                SizedBox(width: 8),
                Text(
                  option.label,
                  style: TextStyle(
                    color: regularFontColor,
                    fontSize: 12,
                    fontWeight: FontWeight.w400,
                  ),
                ),
              ],
            ),
          ),
        ),
        SizedBox(width: 32),
      ]);
    });

    children.removeLast();

    return children;
  }

  @override
  Widget build(BuildContext context) {
    return Row(
      children: getChildren(),
    );
  }
}

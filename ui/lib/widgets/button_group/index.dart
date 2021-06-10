import 'package:flutter/material.dart';

import '../../common/colors.dart';
import 'model.dart';

enum ButtonType {
  error,
  primary,
  defaults,
}

class ButtonGroup extends StatelessWidget {
  final String? selectedKey;
  final List<ButtonGroupItem> buttons;
  final Function? onChange;

  ButtonGroup({
    this.selectedKey,
    required this.buttons,
    this.onChange,
  });

  List<bool> getSelecteds() {
    return [...buttons.map((button) => button.key == selectedKey)];
  }

  List<Widget> getChildren() {
    return [
      ...buttons.map(
        (button) => Container(
          height: 32,
          padding: EdgeInsets.symmetric(horizontal: 12),
          alignment: Alignment.centerLeft,
          child: button.child,
        ),
      )
    ];
  }

  @override
  Widget build(BuildContext context) {
    return SizedBox(
      height: 32,
      child: ToggleButtons(
        selectedColor: Colors.white,
        fillColor: Theme.of(context).primaryColor,
        highlightColor: Theme.of(context).primaryColor,
        focusColor: Colors.blue,
        borderWidth: 1,
        borderRadius: BorderRadius.circular(2),
        borderColor: rgba(228, 235, 241, 1),
        selectedBorderColor: Theme.of(context).primaryColor,
        isSelected: getSelecteds(),
        children: getChildren(),
        onPressed: (index) => onChange!(buttons[index].key),
      ),
    );
  }
}

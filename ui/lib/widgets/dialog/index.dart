import 'package:flutter/material.dart';

import '../../common/colors.dart';

const List<Widget> _defaultLeftButtons = [];

class CommonDialog extends StatelessWidget {
  final String title;
  final double width;
  final Widget content;
  final Widget? actions;
  final List<Widget> buttons;
  final List<Widget> leftButtons;
  final Function onClose;

  CommonDialog({
    required this.title,
    required this.width,
    required this.content,
    this.actions,
    required this.buttons,
    this.leftButtons = _defaultLeftButtons,
    required this.onClose,
  });

  List<Widget> getActions() {
    if (actions != null) {
      return [actions as Widget];
    }
    if (buttons.length == 0 && leftButtons.length == 0) {
      return [];
    }

    return [
      Container(
        width: width,
        height: 56,
        alignment: Alignment.center,
        padding: EdgeInsets.symmetric(
          horizontal: 20,
        ),
        decoration: new BoxDecoration(
          color: Colors.white,
          border: Border(
            top: BorderSide(
              color: defaultDisabledColor,
              width: 1,
            ),
          ),
        ),
        child: Row(
          mainAxisAlignment: MainAxisAlignment.spaceBetween,
          children: [
            Row(children: leftButtons),
            Row(children: buttons),
          ],
        ),
      ),
    ];
  }

  @override
  Widget build(BuildContext context) {
    final size = MediaQuery.of(context).size;
    final contentHeight = size.height - 60 - 60 - 90;

    return AlertDialog(
      title: Container(
        padding: EdgeInsets.only(top: 9, right: 21, bottom: 9, left: 24),
        decoration: new BoxDecoration(
          border: Border(
            bottom: BorderSide(
              color: rgba(226, 232, 240, 1),
              width: 1,
            ),
          ),
        ),
        child: Row(
          children: [
            Expanded(
              child: SelectableText(
                title,
                style: Theme.of(context).textTheme.headline5,
              ),
            ),
            IconButton(
              icon: Icon(Icons.close),
              iconSize: 20,
              padding: EdgeInsets.all(1.0),
              splashRadius: 1.0,
              onPressed: () => onClose(),
            ),
          ],
        ),
      ),
      titlePadding: EdgeInsets.zero,
      content: ConstrainedBox(
        constraints: BoxConstraints(maxHeight: contentHeight),
        child: Scrollbar(
          child: SingleChildScrollView(
            child: content,
          ),
        ),
      ),
      contentPadding: EdgeInsets.all(0),
      actions: getActions(),
      actionsPadding: EdgeInsets.zero,
      buttonPadding: EdgeInsets.zero,
      insetPadding: EdgeInsets.zero,
    );
  }
}

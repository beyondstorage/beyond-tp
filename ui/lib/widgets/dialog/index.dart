import 'package:flutter/material.dart';

class CommonDialog extends StatelessWidget {
  final String title;
  final Widget content;
  final List<Widget> buttons;
  final Function onClose;

  CommonDialog({
    required this.title,
    required this.content,
    required this.buttons,
    required this.onClose
  });

  @override
  Widget build(BuildContext context) {
    final size = MediaQuery.of(context).size;
    final contentHeight = size.height - 60 - 60 - 90;

    return AlertDialog(
      title: Row(
        children: [
          Expanded(
            child: SelectableText(
              title,
              style: Theme.of(context).textTheme.headline4,
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
      titlePadding: EdgeInsets.only(top: 16, right: 20, bottom: 32, left: 20),
      content: ConstrainedBox(
        constraints: BoxConstraints(maxHeight: contentHeight),
        child: Scrollbar(
          child: SingleChildScrollView(
            child: content,
          ),
        ),
      ),
      contentPadding: EdgeInsets.all(0),
      actions: [
        Container(
          width: 600,
          padding: EdgeInsets.symmetric(
            vertical: 12,
            horizontal: 44,
          ),
          decoration: new BoxDecoration(
            color: Colors.white,
            borderRadius: BorderRadius.all(Radius.circular(3.0)),
            boxShadow: [
              BoxShadow(
                offset: Offset(0, -1),
                color: Color.fromRGBO(3, 5, 7, 0.08),
                blurRadius: 3.0,
              )
            ],
          ),
          child: Row(
            mainAxisAlignment: MainAxisAlignment.end,
            children: buttons,
          ),
        ),
      ],
      actionsPadding: EdgeInsets.zero,
      buttonPadding: EdgeInsets.zero,
      insetPadding: EdgeInsets.zero,
    );
  }
}

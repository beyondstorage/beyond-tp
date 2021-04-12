import 'package:flutter/material.dart';

class CommonDialog extends StatelessWidget {
  final String title;
  final Widget content;
  final Widget footer;
  final Function onClose;

  CommonDialog({this.title, this.content, this.footer, this.onClose});

  @override
  Widget build(BuildContext context) {
    return SimpleDialog(
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
      contentPadding: EdgeInsets.all(0),
      children: [
        content,
        footer,
      ],
    );
  }
}

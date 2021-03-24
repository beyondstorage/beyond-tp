import 'package:flutter/material.dart';

class CommonDialog extends StatelessWidget {
  final String title;
  final Widget content;
  final Widget footer;

  CommonDialog({this.title, this.content, this.footer});

  @override
  Widget build(BuildContext context) {
    return SimpleDialog(
      title: Text(
        title,
        style: TextStyle(
          fontSize: 16,
          height: 1.5,
          color: Theme.of(context).textTheme.bodyText1.color,
        ),
      ),
      titlePadding: EdgeInsets.only(top: 16, right: 20, bottom: 32, left: 20),
      contentPadding: EdgeInsets.all(0),
      children: [
        ConstrainedBox(
          constraints: BoxConstraints(maxHeight: 600),
          child: Scrollbar(
            child: SingleChildScrollView(
              child: content,
            ),
          ),
        ),
        footer,
      ],
    );
  }
}

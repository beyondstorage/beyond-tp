import 'package:get/get.dart';
import 'package:flutter/material.dart';

import '../button/index.dart';


class DeleteDialog extends StatelessWidget {
  final String title;
  final Widget child;
  final Function onClose;
  final Function onDelete;

  DeleteDialog({
    this.title,
    this.child,
    this.onClose,
    this.onDelete,
  });

  @override
  Widget build(BuildContext context) {
    return AlertDialog(
      title: Row(
        children: [
          Expanded(
            child: SelectableText(
              title, style: Theme.of(context).textTheme.headline4
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
      titlePadding: EdgeInsets.only(left: 20, top: 16, right: 10, bottom: 16),
      content: SizedOverflowBox(
        alignment: Alignment.topCenter,
        child: child, size: Size(560, 160)
      ),
      actions: <Widget>[
        Button(
          child: Text(
            "Cancel".tr,
            style: Theme.of(context).textTheme.bodyText1,
          ),
          onPressed: () => onClose(),
        ),
        Button(
          type: ButtonType.error,
          child: Text("Delete".tr),
          onPressed: () => onDelete(),
        ),
      ],
      elevation: 24.0,
      actionsPadding: EdgeInsets.symmetric(horizontal: 20, vertical: 12),
    );
  }
}
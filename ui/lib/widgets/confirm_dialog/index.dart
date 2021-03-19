import 'package:get/get.dart';
import 'package:flutter/material.dart';

import '../button/index.dart';


class ConfirmDialog extends StatelessWidget {
  final String title;
  final String content;
  final Function onConfirm;

  ConfirmDialog({
    this.title,
    this.content,
    this.onConfirm,
  });

  @override
  Widget build(BuildContext context) {
    return AlertDialog(
      title: Text(
        title,
        style: Theme.of(context).textTheme.headline4,
      ),
      content: Text(
        content,
        style: Theme.of(context).textTheme.bodyText1,
      ),
      actions: <Widget>[
        Button(
          icon: Icons.delete,
          label: Text("Confirm".tr),
          type: ButtonType.error,
          onPressed: () => onConfirm(),
        ),
      ],
      elevation: 24.0,
    );
  }
}
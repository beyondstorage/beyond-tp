import 'package:get/get.dart';
import 'package:flutter/material.dart';

import '../button/index.dart';
import '../../common/global.dart';

import './title.dart';


class Confirm extends StatelessWidget {
  final String title;
  final String? description;
  final IconData? icon;
  final Function? onClose;
  final Function onConfirm;

  Confirm({
    this.title = "Confirm",
    this.icon = Icons.report_problem,
    this.description,
    this.onClose,
    required this.onConfirm,
  });

  void onClosePressed() {
    Get.back();
    this.onClose!();
  }

  @override
  Widget build(BuildContext context) {
    return AlertDialog(
      title: Align(
        child: IconButton(
          icon: Icon(Icons.close),
          iconSize: 20,
          padding: EdgeInsets.all(1.0),
          splashRadius: 1.0,
          onPressed: () => onClosePressed(),
        ),
        alignment: Alignment.centerRight
      ),
      titlePadding: EdgeInsets.symmetric(vertical: 16, horizontal: 20),
      content: SizedBox(
        width: confirmDialogWidth,
        height: 80.0,
        child: Column(
          children: [
            ConfirmTitle(icon: this.icon, title: this.title),
            Padding(
              padding: EdgeInsets.only(left: 36.0, top: 8.0, bottom: 4.0),
              child: SelectableText(
                this.description!,
                maxLines: 2,
                style: Theme.of(context).textTheme.bodyText2,
              ),
            ),
          ],
        ),
      ),
      contentPadding: EdgeInsets.symmetric(horizontal: 32.0),
      actions: <Widget>[
        Button(
          child: Text(
            "Cancel".tr, style: Theme.of(context).textTheme.bodyText1),
          onPressed: () => onClosePressed(),
        ),
        Button(
          type: ButtonType.error,
          child: Text("Delete".tr),
          onPressed: () => onConfirm(),
        ),
      ],
      elevation: 24.0,
      actionsPadding: EdgeInsets.all(20),
    );
  }
}
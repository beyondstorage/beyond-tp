import 'package:get/get.dart';
import 'package:flutter/material.dart';

import '../button/index.dart';
import '../button/constants.dart';
import '../../common/global.dart';
import './title.dart';

class Confirm extends StatelessWidget {
  final String title;
  final String? description;
  final IconData? icon;
  final Function? onClose;
  final Function onConfirm;
  final Color iconColor;
  final String confirmBtnText;

  Confirm({
      this.title = "Confirm",
      this.icon = const IconData(0xe603, fontFamily: 'tpIcon'),
      this.description,
      this.onClose,
      required this.onConfirm,
      this.iconColor = Colors.redAccent,
      this.confirmBtnText = "Delete",
  });

  void onClosePressed() {
    Get.back();
    this.onClose!();
  }

  get ConfirmBtn {
    List<Widget> horizontal = [
      Row(
        mainAxisAlignment: MainAxisAlignment.end,
        children: [
          Button(
            child: Text("Cancel".tr),
            onPressed: () => onClosePressed(),
          ),
          SizedBox(width: 8,),
          Button(
            type: ButtonType.primary,
            child: Text(
              this.confirmBtnText,
            ),
            onPressed: () => onConfirm(),
          ),
        ],
      )
    ];
    return horizontal;
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
        alignment: Alignment.centerRight,
      ),
      titlePadding: EdgeInsets.symmetric(vertical: 16, horizontal: 20),
      content: SizedBox(
        width: confirmDialogWidth,
        height: 80.0,
        child: Column(
          children: [
            ConfirmTitle(
              icon: this.icon,
              title: this.title,
              color: this.iconColor,
            ),
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
      actions: ConfirmBtn,
      elevation: 24.0,
      actionsPadding: EdgeInsets.all(20),
    );
  }
}

import 'package:get/get.dart';
import 'package:flutter/material.dart';

import '../../../widgets/dialog/index.dart';
import '../../../widgets/button/index.dart';
import 'form.dart';

import 'controller.dart';

class CreateTaskDialog extends GetView<CreateTaskController> {
  final CreateTaskController c = Get.put(CreateTaskController());
  final GlobalKey<FormState> formKey = GlobalKey<FormState>();
  final Function getTasks;

  CreateTaskDialog({this.getTasks});

  void onSubmit() {
    final form = formKey.currentState;

    if (!form.validate()) {
      controller.autoValidateMode(AutovalidateMode.always);
    } else {
      form.save();
      controller.onSubmit(getTasks);
    }
  }

  @override
  Widget build(BuildContext context) {
    return CommonDialog(
      title: 'Create task'.tr,
      content: Container(
        width: 600.0,
        child: Padding(
          padding: EdgeInsets.symmetric(horizontal: 20.0),
          child: CreateTaskForm(formKey, onSubmit),
        ),
      ),
      footer: Container(
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
          children: [
            Button(
              child: Text(
                "Cancel".tr,
                style: Theme.of(context).textTheme.bodyText1,
              ),
              type: ButtonType.defaults,
              onPressed: () => Get.back(),
            ),
            SizedBox(width: 12),
            Button(
              child: Text("Submit".tr),
              type: ButtonType.primary,
              onPressed: onSubmit,
            ),
          ],
        ),
      ),
    );
  }
}

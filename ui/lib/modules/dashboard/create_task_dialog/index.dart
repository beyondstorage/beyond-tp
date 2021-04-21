import 'package:get/get.dart';
import 'package:flutter/material.dart';

import '../../../widgets/dialog/index.dart';
import '../../../widgets/button/index.dart';
import 'form.dart';

import 'controller.dart';

class CreateTaskDialog extends StatelessWidget {
  final CreateTaskController controller = Get.put(CreateTaskController());
  final GlobalKey<FormState> formKey = GlobalKey<FormState>();
  final Function getTasks;

  CreateTaskDialog({ required this.getTasks });

  void onSubmit() {
    final form = formKey.currentState;

    if (!form!.validate()) {
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
        width: 600,
        child: Padding(
          padding: EdgeInsets.symmetric(horizontal: 20.0),
          child: CreateTaskForm(formKey, onSubmit),
        ),
      ),
      buttons: [
        Button(
          child: Text(
            "Cancel".tr,
            style: Theme.of(context).textTheme.bodyText1,
          ),
          type: ButtonType.defaults,
          onPressed: controller.closeDialog,
        ),
        SizedBox(width: 12),
        Button(
          child: Text("Submit".tr),
          type: ButtonType.primary,
          onPressed: onSubmit,
        ),
      ],
      onClose: controller.closeDialog,
    );
  }
}

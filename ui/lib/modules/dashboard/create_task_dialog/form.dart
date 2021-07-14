import 'package:get/get.dart';
import 'package:flutter/material.dart';
import 'package:form_validator/form_validator.dart';

import 'source_library_setting.dart';
import 'target_library_setting.dart';
import 'controller.dart';

class CreateTaskForm extends GetView<CreateTaskController> {
  final GlobalKey<FormState> sourceFormKey;
  final GlobalKey<FormState> targetFormKey;
  final Function onSubmit;

  CreateTaskForm(this.sourceFormKey, this.targetFormKey, this.onSubmit);

  @override
  Widget build(BuildContext context) {
    return Container(
      width: 540,
      alignment: Alignment.topLeft,
      child: Obx(
        () => controller.step.value == 1
            ? SourceLibrarySetting(sourceFormKey)
            : TargetLibrarySetting(targetFormKey),
      ),
    );
  }
}

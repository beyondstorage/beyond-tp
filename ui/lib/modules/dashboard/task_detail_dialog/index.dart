import 'package:get/get.dart';
import 'package:flutter/material.dart';

import '../../../models/task.dart';
import '../../../widgets/dialog/index.dart';
import '../../../widgets/button/index.dart';
import './task_detail.dart';
import '../controller.dart';

class TaskDetailDialog extends GetView<DashboardController> {
  final String taskId;

  TaskDetailDialog({this.taskId}) {
    controller.getTaskDetail(taskId);
  }

  void closeDialog() {
    Get.back();
    controller.taskDetail(TaskDetail.fromMap({}));
  }

  @override
  Widget build(BuildContext context) {
    return CommonDialog(
      title: 'Task detail'.tr,
      content: Container(
        width: 600,
        child: Padding(
          padding: EdgeInsets.symmetric(horizontal: 20),
          child: TaskDetailForm(),
        ),
      ),
      buttons: [
        Button(
          child: Text("Finish".tr),
          type: ButtonType.primary,
          onPressed: closeDialog,
        ),
      ],
      onClose: closeDialog,
    );
  }
}

import 'package:get/get.dart';
import 'package:flutter/material.dart';
import 'package:ui/modules/dashboard/task_detail/return_task_list.dart';

import 'task_description.dart';
import 'task_information.dart';

import '../controller.dart';

class TaskDetail extends GetView<DashboardController> {

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        ReturnToList(
          title: "Return tasks tist".tr,
          onTap: () {
            controller.showDetail(false);
            controller.detailTaskId('');
          }
        ),
        TaskDescription(),
        TaskInformation()
      ],
    );
  }
}

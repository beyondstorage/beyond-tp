import 'package:get/get.dart';
import 'package:flutter/material.dart';
import 'package:ui/modules/dashboard/task_detail/index.dart';

import '../layout/index.dart';
import '../../widgets/page_description/index.dart';
import '../../widgets/empty_entry_list/index.dart';

import 'entry_list.dart';
import 'create_task_dialog/index.dart';
import 'controller.dart';

class Dashboard extends GetView<DashboardController> {
  final DashboardController c = Get.put(DashboardController());

  @override
  Widget build(BuildContext context) {
    c.getTasks();

    return Layout(
      child: Obx(() => controller.showDetail.value
        ? TaskDetail()
        : Column(
            children: [
              PageDescription(
                icon: IconData(0xe600, fontFamily: 'tpIcon'),
                title: 'Tasks'.tr,
                subtitle: "Create and manage your data migration tasks".tr,
              ),
              controller.tasks.value.length() == 0
                ? EmptyEntryList(
                  icon: IconData(0xe600, fontFamily: 'tpIcon'),
                  title: 'The task list is empty'.tr,
                  subTitle: 'Please click the button below to create a task'.tr,
                  buttonText: 'Create task'.tr,
                  onClick: () => Get.dialog(CreateTaskDialog(
                    name: 'DM task 1',
                    getTasks: controller.getTasks,
                  )),
                )
                : EntryList(),
            ],
        )),
    );
  }
}

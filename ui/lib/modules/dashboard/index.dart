import 'package:get/get.dart';
import 'package:flutter/material.dart';

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
      child: Column(
        children: [
          PageDescription(
            icon: Icons.source,
            title: 'Tasks'.tr,
            subtitle: "Create and manage your data migration tasks".tr,
          ),
          Obx(
            () => controller.tasks.value.length() == 0
                ? EmptyEntryList(
                    icon: Icons.source,
                    title: 'The task list is empty'.tr,
                    subTitle:
                        'Please click the button below to create a task'.tr,
                    buttonText: 'Create task'.tr,
                    onClick: () => Get.dialog(CreateTaskDialog(
                      name: 'DM Task 1',
                      getTasks: controller.getTasks,
                    )),
                  )
                : EntryList(),
          ),
        ],
      ),
    );
  }
}

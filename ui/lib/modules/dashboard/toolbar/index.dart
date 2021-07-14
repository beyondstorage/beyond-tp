import 'package:get/get.dart';
import 'package:flutter/material.dart';

import '../../../common/colors.dart';
import '../../../widgets/button/index.dart';
import '../../../widgets/button/constants.dart';
import '../../../widgets/search_input/index.dart';

import 'filter.dart';
import '../controller.dart';
import '../create_task_dialog/index.dart';

class Operation {
  final String key;
  final String value;

  Operation({
    required this.key,
    required this.value,
  });
}

List<Operation> operations = [
  Operation(key: 'all', value: 'All'.tr),
  Operation(key: 'running', value: 'Running'.tr),
  Operation(key: 'not_running', value: 'Not Running'.tr),
  Operation(key: 'completed', value: 'Completed'.tr),
];

class Toolbar extends GetView<DashboardController> {
  void onPressed(String op) {
    controller.getTasks();
  }

  void handleCreateTask() {
    final taskNum = controller.tasks.value.length();
    final newTaskName = 'DM Task ${taskNum + 1}';
    Get.dialog(
        CreateTaskDialog(name: newTaskName, getTasks: controller.getTasks));
  }

  @override
  Widget build(BuildContext context) {
    String key = 'all';
    int total = controller.tasks.value.toList().length;

    return Container(
      margin: EdgeInsets.only(bottom: 16.0),
      child: Column(
        children: [
          Container(
            margin: EdgeInsets.only(bottom: 24.0),
            decoration: BoxDecoration(
              color: Colors.white,
              boxShadow: [
                BoxShadow(offset: Offset(0, 1), color: rgba(226, 232, 240, 1)),
                BoxShadow(offset: Offset(-1, 0), color: Colors.white),
                BoxShadow(offset: Offset(1, 0), color: Colors.white),
              ],
            ),
            child: Row(
              children: operations
                  .map((Operation op) => Filter(
                        title: op.value,
                        selected: key == op.key,
                        counts: op.key == 'all' ? total : 0,
                        onPressed: () => onPressed(op.key),
                      ))
                  .toList(),
            ),
          ),
          Row(
            children: [
              Button(
                icon: Icons.add,
                child: Text("Create task".tr),
                type: ButtonType.primary,
                onPressed: handleCreateTask,
              ),
              Expanded(child: Text('')),
              Obx(() => SearchInput(
                    defaultValue: controller.filters.value,
                    onClear: () => controller.filters(''),
                    onChange: (text) => controller.filters(text),
                    placeholder: 'Search for Task Name / ID / Library Type'.tr,
                  )),
            ],
          )
        ],
      ),
    );
  }
}

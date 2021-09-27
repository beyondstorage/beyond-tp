import 'package:get/get.dart';
import 'package:flutter/material.dart';

import 'controller.dart';

import '../../common/colors.dart';
import '../../widgets/more_actions/index.dart';
import '../../widgets/confirm/index.dart';

import './task_detail_dialog/index.dart';

class EntryActions extends GetView<DashboardController> {
  final Map<String, dynamic> data;

  EntryActions({ required this.data });

  @override
  Widget build(BuildContext context) {
    String status = data["status"];
    final startAble = status == "Created" || status == "Stopped";

    return Container(
      padding: EdgeInsets.symmetric(vertical: 10, horizontal: 16),
      child: SingleChildScrollView(
        scrollDirection: Axis.horizontal,
        child: Row(
          children: [
            TextButton(
              child: Padding(
                padding: EdgeInsets.symmetric(horizontal: 8),
                child: Text(
                  "View detail".tr,
                  style: TextStyle(
                    fontSize: 12,
                    color: rgba(3, 102, 214, 1),
                  ),
                ),
              ),
              onPressed: () {
                // Get.dialog(TaskDetailDialog(taskId: data["id"]));
                controller.showDetail(true);
              },
            ),
            TextButton(
              child: Padding(
                padding: EdgeInsets.symmetric(horizontal: 8),
                child: Text(
                  "Start".tr,
                  style: TextStyle(
                    fontSize: 12,
                    color: rgba(3, 102, 214, startAble ? 1 : .5),
                  ),
                ),
              ),
              onPressed:
              startAble ? () => controller.runTask(data["id"]) : null,
            ),
          ],
        ),
      ),
    );
  }
}

class EntryMoreActions extends GetView<DashboardController> {
  final dynamic value;
  final Map<String, dynamic> data;

  EntryMoreActions({this.value, required this.data});

  @override
  Widget build(BuildContext context) {

    return MoreActions(
      onSelected: (String op) {
        Get.dialog(
            Confirm(
                title: "Are you sure to delete this task?".tr,
                description: "This task has been completed and will no longer be displayed in the task list after deletion.".tr,
                onConfirm: () {
                  controller.deleteTask(data["id"]).then((result) {
                    Get.back();
                  });
                }
            )
        );
      },
      itemBuilder: (BuildContext context) => [
        PopupMenuItem(
          value: "delete",
          height: 32.0,
          child: Text("Delete task".tr,
              style: TextStyle(
                fontSize: 12.0,
                color: Theme.of(context).errorColor,
              )),
        ),
      ],
    );
  }
}

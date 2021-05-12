import 'package:get/get.dart';
import 'package:flutter/material.dart';
// ignore: import_of_legacy_library_into_null_safe
import 'package:flutter_layout_grid/flutter_layout_grid.dart';

import 'controller.dart';

import '../../common/global.dart';
import '../../common/colors.dart';
import '../../widgets/grid_table/model.dart';
import '../../widgets/grid_table/index.dart';
import '../../widgets/more_actions/index.dart';
import '../../widgets/confirm/index.dart';

import './task_status.dart';
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
              onPressed: () => Get.dialog(TaskDetailDialog(taskId: data["id"])),
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
    // String name = data["name"];

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

class EntryList extends GetView<DashboardController> {
  final List<GridTableCol> columns = [
    GridTableCol(title: "Name".tr, dataIndex: "name"),
    GridTableCol(
      title: "Status".tr,
      dataIndex: "status",
      render: (value, data) => TaskStatus(value),
    ),
    GridTableCol(title: "Created at".tr, dataIndex: "createdAt"),
    GridTableCol(title: "Updated at".tr, dataIndex: "updatedAt"),
    GridTableCol(
      title: "Actions",
      dataIndex: 'actions',
      render: (value, data) => EntryActions(data: data),
    ),
    GridTableCol(
      width: 40.px,
      dataIndex: 'moreActions',
      render: (value, data) => EntryMoreActions(value: value, data: data),
    ),
  ];

  @override
  Widget build(BuildContext context) {
    return Obx(() => GridTable(
          columns: columns,
          dataList: controller.tasks.value.toList(),
          maxHeight: Get.height - globalHeaderHeight - 180.0,
        ));
  }
}

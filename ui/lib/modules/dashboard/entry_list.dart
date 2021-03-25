import 'package:get/get.dart';
import 'package:flutter/material.dart';
import 'package:flutter_layout_grid/flutter_layout_grid.dart';
import 'package:ui/widgets/delete_dialog/index.dart';

import 'controller.dart';

import '../../common/global.dart';
import '../../widgets/grid_table/model.dart';
import '../../widgets/grid_table/index.dart';
import '../../widgets/more_actions/index.dart';
import '../../widgets/delete_dialog/index.dart';

class EntryActions extends GetView<DashboardController> {
  final dynamic value;
  final Map<String, dynamic> data;

  EntryActions({ this.value, this.data });

  @override
  Widget build(BuildContext context) {
    String name = data["name"];

    return MoreActions(
      onSelected: (String op) {
        showDialog(
          context: context,
          builder: (BuildContext context) => DeleteDialog(
            title: "Delete Task".tr,
            child: SelectableText(
              "Confirm to delete task $name?",
              style: Theme.of(context).textTheme.bodyText1,
            ),
            onDelete: () {
              controller.deleteTask(data["id"]).then((result) {
                Navigator.pop(context, true);
              });
            },
            onClose: () => Navigator.pop(context, true),
          ),
        );
        // controller.deleteTask(data["id"]);
      },
      itemBuilder: (BuildContext context) => [
        PopupMenuItem(
          value: "delete",
          height: 32.0,
          child: Text(
            "Delete task".tr,
            style: TextStyle(
              fontSize: 12.0,
              color: Theme.of(context).errorColor,
            )
          ),
        ),
      ],
    );
  }
}

class EntryList extends GetView<DashboardController> {
  final List<GridTableCol> columns = [
    GridTableCol(title: "Name".tr, dataIndex: "name"),
    GridTableCol(title: "Status".tr, dataIndex: "status"),
    GridTableCol(title: "Created at".tr, dataIndex: "createdAt"),
    GridTableCol(title: "Updated at".tr, dataIndex: "updatedAt"),
    GridTableCol(
      width: 40.px,
      dataIndex: 'actions',
      render: (value, data) => EntryActions(value: value, data: data),
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

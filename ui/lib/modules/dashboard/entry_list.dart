import 'package:get/get.dart';
import 'package:flutter/material.dart';
import 'package:flutter_layout_grid/flutter_layout_grid.dart';

import '../../common/global.dart';
import '../../widgets/grid_table/model.dart';
import '../../widgets/grid_table/index.dart';
import '../../widgets/page_container/index.dart';

import 'toolbar/index.dart';
import 'controller.dart';
import 'task_status.dart';
import 'entry_actions.dart';


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
    return WidgetContainer(
      child: Padding(
        padding: EdgeInsets.symmetric(vertical: 12.0, horizontal: 20.0),
        child: Column(
          children: [
            Toolbar(),
            Obx(() => GridTable(
              columns: columns,
              dataList: controller.tasks.value.toList(),
              maxHeight: Get.height - globalHeaderHeight - 180.0,
            ))
          ],
        ),
      ),
    );
  }
}

import 'package:get/get.dart';
import 'package:flutter/material.dart';

import 'controller.dart';

import '../../common/global.dart';
import '../../widgets/grid_table/model.dart';
import '../../widgets/grid_table/index.dart';

class EntryList extends GetView<DashboardController> {
  final List<GridTableCol> columns = [
    GridTableCol(title: "Name", dataIndex: "name"),
    GridTableCol(title: "Status", dataIndex: "status"),
    GridTableCol(title: "Created at", dataIndex: "createdAt"),
    GridTableCol(title: "Updated at", dataIndex: "updatedAt"),
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

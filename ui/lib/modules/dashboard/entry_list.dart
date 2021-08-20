import 'package:get/get.dart';
import 'package:flutter/material.dart';
import 'package:flutter_layout_grid/flutter_layout_grid.dart';
import 'package:date_format/date_format.dart';

import '../../common/global.dart';
import '../../common/colors.dart';
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
    GridTableCol(
        title: "Creation time".tr,
        dataIndex: "createdAt",
        render: (value, data) {
          return Container(
            padding: EdgeInsets.symmetric(vertical: 12.0, horizontal: 16.0),
            child: SelectableText(
              formatDate(
                DateTime.now(),
                [yyyy, "-", mm, "-", dd, " ", HH, ":", nn, ":", ss],
              ),
              style: TextStyle(
                fontSize: 12,
                color: regularFontColor,
              ),
            ),
          );
        }),
    GridTableCol(
      title: "Operation",
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
            Obx(
              () => GridTable(
                columns: columns,
                dataList: controller.tasks.value.toList(),
                maxHeight: Get.height - 336.0,
              ),
            )
          ],
        ),
      ),
    );
  }
}

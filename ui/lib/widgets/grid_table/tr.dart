import 'package:flutter/gestures.dart';
import 'package:flutter/rendering.dart';
import 'package:get/get.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter_layout_grid/flutter_layout_grid.dart';

import 'td.dart';
import 'model.dart';
import '../../common/colors.dart';

class GridTableTR extends StatelessWidget {
  final List<TrackSize> columnSizes;
  final List<GridTableCol> columns;
  final Map<String, dynamic> data;

  final RxBool hovered = false.obs;

  Color get color {
    if (hovered.value) {
      return rgba(245, 247, 250, 0.5);
    }

    return Colors.white;
  }

  GridTableTR({
    required this.columns,
    required this.columnSizes,
    required this.data,
  });

  void onHovered(PointerEnterEvent event) {
    hovered(true);
  }

  void onExit(PointerExitEvent event) {
    hovered(false);
  }

  @override
  Widget build(BuildContext context) {
    return MouseRegion(
      onExit: onExit,
      onEnter: onHovered,
      // cursor: SystemMouseCursors.click,
      child: Obx(() => Container(
        decoration: BoxDecoration(
          color: color,
          border: Border(
            bottom: BorderSide(
              width: 1.0,
              color: rgba(228, 235, 241, 1),
            ),
          )
        ),
        child: LayoutGrid(
          rowSizes: [44.px],
          columnSizes: columnSizes,
          children: columns.map((col) {
            if (col.render != null) {
              return col.render!(data[col.dataIndex], data);
            }

            return GridTableTD(
              value: data[col.dataIndex],
              style: Theme.of(context).dataTableTheme.dataTextStyle,
            );
          }).toList()
        ),
      )),
    );
  }
}
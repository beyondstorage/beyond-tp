import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter_layout_grid/flutter_layout_grid.dart';

import 'td.dart';
import 'model.dart';
import '../../common/colors.dart';

class GridTableHeader extends StatelessWidget {
  final List<TrackSize> columnSizes;
  final List<GridTableCol> columns;

  GridTableHeader({ this.columns, this.columnSizes });

  @override
  Widget build(BuildContext context) {
    return Container(
      decoration: BoxDecoration(
        color: rgba(245, 247, 250, 1),
        border: Border(
          bottom: BorderSide(
            color: rgba(228, 235, 241, 1),
            width: 1.0
          )
        )
      ),
      child: LayoutGrid(
        rowSizes: [44.px],
        columnSizes: columnSizes,
        children: columns.map((col) => GridTableTD(
          value: col.title,
          style: Theme.of(context).dataTableTheme.headingTextStyle,
        )).toList()
      ),
    );
  }
}
import 'package:flutter/material.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter_layout_grid/flutter_layout_grid.dart';

import 'tr.dart';
import 'thead.dart';
import 'model.dart';

class GridTable extends StatelessWidget {
  final double? maxHeight;
  final List<GridTableCol> columns;
  final List<Map<String, dynamic>> dataList;

  GridTable({
    this.maxHeight,
    required this.dataList,
    required this.columns,
  });

  List<TrackSize> get columnSizes {
    return columns.map((col) => col.width ?? 1.fr).toList();
  }

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        GridTableHeader(columns: columns, columnSizes: columnSizes),
        Scrollbar(
          child: LayoutBuilder(builder: (context, constrints) => ConstrainedBox(
            child: ListView(
              shrinkWrap: true,
              children: dataList.map((data) => GridTableTR(
                data: data,
                columns: columns,
                columnSizes: columnSizes,
              )).toList(),
            ),
            constraints: BoxConstraints(
              maxWidth: constrints.maxWidth,
              maxHeight: maxHeight ?? constrints.maxHeight,
            ),
          ))
        ),
      ],
    );
  }
}
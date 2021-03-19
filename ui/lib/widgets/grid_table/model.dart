import 'package:flutter/cupertino.dart';
import 'package:flutter_layout_grid/flutter_layout_grid.dart';

class GridTableCol {
  String title;
  @required String dataIndex;
  TrackSize width;

  GridTableCol({
    this.title,
    this.width,
    this.dataIndex,
  });
}

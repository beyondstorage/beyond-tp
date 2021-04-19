import 'package:flutter/cupertino.dart';
import 'package:flutter_layout_grid/flutter_layout_grid.dart';

class GridTableCol {
  String? title;
  String dataIndex;
  TrackSize? width;
  Widget Function(dynamic value, Map<String, dynamic> obj)? render;

  GridTableCol({
    this.title,
    this.width,
    required this.dataIndex,
    this.render,
  });
}

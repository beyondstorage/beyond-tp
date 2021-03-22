import 'package:flutter/material.dart';

class GridTableTD extends StatelessWidget {
  final dynamic value;
  final TextStyle style;

  GridTableTD({ this.value, this.style });

  @override
  Widget build(BuildContext context) {
    return Container(
      padding: EdgeInsets.symmetric(vertical: 12.0, horizontal: 16.0),
      child: Text(
        value,
        style: style,
        maxLines: 1,
        softWrap: false,
        overflow: TextOverflow.ellipsis,
      ),
    );
  }
}
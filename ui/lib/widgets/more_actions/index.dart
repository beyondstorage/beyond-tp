import 'package:flutter/material.dart';

class MoreActions extends StatelessWidget {
  final PopupMenuItemSelected<String> onSelected;
  final PopupMenuItemBuilder<String> itemBuilder;

  MoreActions({ this.onSelected, this.itemBuilder });

  @override
  Widget build(BuildContext context) {
    return PopupMenuButton(
      onSelected: onSelected,
      itemBuilder: itemBuilder,
      // itemBuilder: (BuildContext context) => [
      //   PopupMenuItem(
      //     value: "delete",
      //     child: Text("Delete".tr),
      //   ),
      // ],
      icon: Icon(Icons.more_vert, size: 18.0),
    );
  }
}
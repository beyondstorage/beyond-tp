import 'package:get/get.dart';
import 'package:flutter/material.dart';

class MoreActions extends StatelessWidget {
  final PopupMenuItemSelected<String> onSelected;
  final PopupMenuItemBuilder<String> itemBuilder;
  final bool enabled;

  MoreActions({
    required this.onSelected,
    required this.itemBuilder,
    this.enabled = true,
  });

  @override
  Widget build(BuildContext context) {
    return PopupMenuButton(
      onSelected: onSelected,
      itemBuilder: itemBuilder,
      icon: Icon(Icons.more_vert),
      iconSize: 18.0,
      enabled: enabled,
      tooltip: "More actions".tr,
      offset: Offset(0, 36)
      // itemBuilder: (BuildContext context) => [
      //   PopupMenuItem(
      //     value: "delete",
      //     child: Text("Delete".tr),
      //   ),
      // ],
    );
  }
}
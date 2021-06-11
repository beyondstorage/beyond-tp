import 'package:flutter/material.dart';

import '../../common/colors.dart';

class SearchInput extends StatelessWidget {
  final String? placeholder;
  final String? defaultValue;
  final bool disabled;
  final double width;
  final void Function(String text) onChange;
  final VoidCallback onClear;

  SearchInput({
    this.placeholder,
    this.defaultValue = '',
    this.width = 320.0,
    this.disabled = false,
    required this.onClear,
    required this.onChange,
  });

  @override
  Widget build(BuildContext context) {
    TextEditingController c = TextEditingController.fromValue(
      TextEditingValue(
        text: defaultValue ?? '',
        selection: TextSelection.fromPosition(
          TextPosition(
            affinity: TextAffinity.downstream, offset: defaultValue!.length
          )
        )
      )
    );

    return Container(
      width: width,
      height: 32.0,
      child: TextField(
        controller: c,
        style: Theme.of(context).textTheme.bodyText1,
        decoration: InputDecoration(
          isDense: true,
          prefixIcon: Icon(Icons.search, size: 16.0),
          suffixIcon: c.text.length > 0 ? IconButton(
            icon: Icon(Icons.cancel, size: 16.0),
            onPressed: () {
              c.text = '';
              onClear();
            },
          ) : null,
          contentPadding: EdgeInsets.symmetric(horizontal: 0.0, vertical: 0.0),
          border: OutlineInputBorder(
            borderSide: BorderSide(color: defaultColor),
            borderRadius: BorderRadius.circular(18.0),
          ),
          hintText: placeholder ?? '',
        ),
        onEditingComplete: () => onChange(c.text),
      ),
    );
  }
}

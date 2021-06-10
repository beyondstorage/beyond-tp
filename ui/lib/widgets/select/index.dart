import 'package:flutter/material.dart';

import './model.dart';

class Select extends StatelessWidget {
  final String? hint;
  final String? value;
  final String? Function(String?)? validator;
  final List<SelectOption> options;
  final void Function(String?)? onChange;
  final void Function(String?)? onSaved;

  Select({
    this.hint,
    this.value,
    this.validator,
    required this.options,
    this.onChange,
    this.onSaved,
  });

  @override
  Widget build(BuildContext context) {
    return Expanded(
      child: DropdownButtonFormField(
        style: TextStyle(fontSize: 12),
        decoration: InputDecoration(
          border: const OutlineInputBorder(),
          isDense: true,
          contentPadding: EdgeInsets.symmetric(horizontal: 12, vertical: 10),
        ),
        hint: hint == null ? null : Text(hint as String),
        value: value,
        validator: validator,
        onChanged: onChange == null ? null : onChange,
        onSaved: onSaved == null ? null : onSaved,
        items: [
          ...options.map((option) => DropdownMenuItem(
                value: option.value,
                child: Text(
                  option.label,
                  style: TextStyle(
                    color: Theme.of(context).primaryColorLight,
                  ),
                ),
              ))
        ],
      ),
    );
  }
}

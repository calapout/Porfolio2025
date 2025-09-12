import type { Schema, Struct } from '@strapi/strapi';

export interface UtilRealizationPeriod extends Struct.ComponentSchema {
  collectionName: 'components_util_realization_periods';
  info: {
    displayName: 'Realization Period';
    icon: 'calendar';
  };
  attributes: {
    EndingYear: Schema.Attribute.Integer;
    StartingYear: Schema.Attribute.Integer &
      Schema.Attribute.Required &
      Schema.Attribute.DefaultTo<2025>;
  };
}

declare module '@strapi/strapi' {
  export module Public {
    export interface ComponentSchemas {
      'util.realization-period': UtilRealizationPeriod;
    }
  }
}

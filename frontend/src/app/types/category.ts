export type CategoryType = 'income' | 'expense';

export type Category = {
  id: number;
  type: CategoryType;
  name: string;
  created_at?: string;
  updated_at?: string;
};

export type CategoriesResponse = {
  categories: Category[];
};

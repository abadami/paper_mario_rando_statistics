import type { CategoryDetails } from "../types";

export async function fetchCategoryData(): Promise<CategoryDetails> {
  try {
    const response = await fetch("https://racetime.gg/pm64r/data");

    return response.json();
  } catch (error) {
    console.log(error);
    return {
      name: "",
      short_name: "",
      slug: "",
      url: "",
      data_url: "",
      image: "",
      info: "",
      streaming_required: false,
      goals: [],
    };
  }
}

export async function fetchCategoryGoals(): Promise<string[]> {
  try {
    const categoryData: CategoryDetails = await fetchCategoryData();

    return categoryData.goals;
  } catch (error) {
    console.log(error);
    return [];
  }
}

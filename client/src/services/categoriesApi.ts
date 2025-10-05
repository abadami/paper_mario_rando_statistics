import type { CategoryDetails } from "../types";
import { API_CONFIG } from "../config";

export async function fetchCategoryData(): Promise<CategoryDetails> {
  try {
    const response = await fetch(API_CONFIG.ENDPOINTS.CATEGORY_DATA);

    if (!response.ok) {
      throw new Error("Failed to fetch category data");
    }

    return response.json();
  } catch (error) {
    console.error("Error fetching category data:", error);
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
    const categoryData = await fetchCategoryData();
    return categoryData.goals;
  } catch (error) {
    console.error("Error fetching category goals:", error);
    return [];
  }
}

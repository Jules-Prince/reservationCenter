import axios from 'axios';

const API_URL = 'http://localhost:8080';

// Fetch all availabilities
export const getAvailabilities = async () => {
  try {
    const response = await axios.get(`${API_URL}/availabilities`);
    return response.data;
  } catch (error) {
    console.error('Error fetching availabilities:', error);
    throw error;
  }
};

// Add a new availability (example)
export const createAvailability = async (data: any) => {
  try {
    const response = await axios.post(`${API_URL}/availabilities`, data);
    return response.data;
  } catch (error) {
    console.error('Error creating availability:', error);
    throw error;
  }
};

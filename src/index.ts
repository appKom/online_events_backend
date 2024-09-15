import express, { Request, Response } from "express";
import { createClient } from "@supabase/supabase-js";
import dotenv from "dotenv";

dotenv.config();

const app = express();
const port = 3000;

const supabaseUrl = process.env.SUPABASE_URL;
const supabaseKey = process.env.SUPABASE_KEY;

if (!supabaseUrl || !supabaseKey) {
  throw new Error("Please provide SUPABASE_URL and SUPABASE_KEY");
}

const supabase = createClient(supabaseUrl, supabaseKey);

app.use(express.json());

app.get("/events", async (req: Request, res: Response) => {
  const { data, error } = await supabase.from("events").select("*");

  if (error) {
    return res.status(500).json({ error: error.message });
  }

  res.json(data);
});

app.listen(port, () => {
  console.log(`Server is running on http://localhost:${port}`);
});

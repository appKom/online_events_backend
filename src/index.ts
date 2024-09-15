import express, { Request, Response } from "express";
import { createClient } from "@supabase/supabase-js";

const app = express();
const port = 3000;

const supabaseUrl = "https://your-supabase-url.supabase.co";
const supabaseKey = "your-anon-key";
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

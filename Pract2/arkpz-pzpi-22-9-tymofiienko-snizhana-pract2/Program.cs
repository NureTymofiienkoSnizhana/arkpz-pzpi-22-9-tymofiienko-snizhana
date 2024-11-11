using System;
using System.Collections.Generic;
using System.Data.SqlClient;
using System.Diagnostics;

namespace pz2_code_example
{
    internal class Program
    {
        private double price = 100;
        private double total;
        private string platform = "MAC";
        private string browser = "IE";    
        private int resize = 1;          
        private bool wasInitialized = true; 

        static void Main(string[] args)
        {
            
        }
        bool IsSpecialDeal()
        {
            return true; 
        }
        void Send()
        {
            Console.WriteLine("Send called. Total: " + total);
        }

        // Код до рефакторингу
        void RenderBanner()
        {
            if ((platform.ToUpper().IndexOf("mac") > -1) &&
                (browser.ToUpper().IndexOf("ie") > -1) &&
                wasInitialized && resize > 0)
            {
                // do something
            }
        }

        // Код після рефакторингу
        void RenderBanner2()
        {
            bool isMacOs = platform.ToUpper().IndexOf("MAC") > -1;
            bool isIE = browser.ToUpper().IndexOf("IE") > -1;
            bool wasResized = resize > 0;

            if (isMacOs && isIE && wasInitialized && wasResized)
            {
                // do something
            }
        }

        // Код до рефакторингу
        string FoundPerson(string[] people)
        {
            for (int i = 0; i < people.Length; i++)
            {
                if (people[i].Equals("Don"))
                {
                    return "Don";
                }
                if (people[i].Equals("John"))
                {
                    return "John";
                }
                if (people[i].Equals("Kent"))
                {
                    return "Kent";
                }
            }
            return String.Empty;
        }

        // Код після рефакторингу
        string FoundPerson2(string[] people)
        {
            List<string> candidates = new List<string>() { "Don", "John", "Kent" };
            for (int i = 0; i < people.Length; i++)
            {
                if (candidates.Contains(people[i]))
                {
                    return people[i];
                }
            }
            return String.Empty;
        }

        // Код до рефакторингу
        void Deal()
        {
            if (IsSpecialDeal())
            {
                total = price * 0.95;
                Send();
            }
            else
            {
                total = price * 0.98;
                Send();
            }
        }

        // Код після рефакторингу
        void Deal2()
        {
            if (IsSpecialDeal())
            {
                total = price * 0.95;
            }
            else
            {
                total = price * 0.98;
            }
            Send();
        }

        // Власний код до рефакторингу
        private void Search(DataGridView dgw)
        {
            dgw.Rows.Clear();
            if (selectedTable == "Motorcycles")
            {
                string searchQuery = $"SELECT Motorcycles.motorcycle_id" + textBox_search.Text + "%'";

                SqlCommand com = new SqlCommand(searchQuery, motostoreDB.getConnection());
                motostoreDB.openConnection();
                SqlDataReader read = com.ExecuteReader();

                while (read.Read())
                {
                    ReadSingleRowMoto(dgw, read);
                }
                read.Close();
            }
            else if (selectedTable == "Orders")
            {
                string searchQuery = $"SELECT Orders.order_id " + textBox_search.Text + "%'";

                SqlCommand com = new SqlCommand(searchQuery, motostoreDB.getConnection());
                motostoreDB.openConnection();
                SqlDataReader read = com.ExecuteReader();

                while (read.Read())
                {
                    ReadSingleRowOrders(dgw, read);
                }
                read.Close();
            }
        }


        // Власний код після рефакторингу
        private void Search(DataGridView dgw)
        {
            string searchQuery = "";
            dgw.Rows.Clear();

            if (selectedTable == "Motorcycles")
            {
                searchQuery = $"SELECT Motorcycles.motorcycle_id" + textBox_search.Text + "%'";
            }
            else if (selectedTable == "Orders")
            {
                searchQuery = $"SELECT Orders.order_id " + textBox_search.Text + "%'";
            }

            SqlCommand com = new SqlCommand(searchQuery, motostoreDB.getConnection());
            motostoreDB.openConnection();
            SqlDataReader read = com.ExecuteReader();

            while (read.Read())
            {
                ReadSingleRowMoto(dgw, read);
            }

            read.Close();
        }
    }
}

library(ggplot2)
library(ggthemes)

l = list()
l$n = c(1:100)
l$p = 2/(l$n + 2)

l <- as.data.frame(l)

title_text = "Probability assigned by Laplace's rule of succession\nas the number of trials increases"
label_x_axis = "number of trials"
label_y_axis = "probability"
ggplot(data=l, aes(x=n, y=p))+
  geom_point(size = 0.5, color="navyblue")+
  labs(
    title=title_text,
    subtitle=element_blank(),
    x=label_x_axis, 
    y=label_y_axis 
  ) +
  theme_tufte() + 
  theme(
    legend.title = element_blank(),
    plot.title = element_text(hjust = 0.5),
    plot.subtitle = element_text(hjust = 0.5), 
    legend.position="bottom",
    legend.box="vertical",
    axis.text.x=element_text(angle=60, hjust=1),
    plot.background=element_rect(fill = "white",colour = NA)
  )

